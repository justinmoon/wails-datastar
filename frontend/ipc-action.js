// ipc-action.js - Datastar plugin for Wails IPC integration
// import { dispatchSSE } from 'datastar';

export function dispatchSSE(
  type,
  elId,
  argsRaw,
) {
  document.dispatchEvent(
    new CustomEvent(DATASTAR_SSE_EVENT, {
      detail: { type, elId, argsRaw },
    }),
  )
}


// Constants for Datastar event lifecycle
const DATASTAR_SSE_EVENT = `datastar-sse`
const STARTED = 'datastar-started';
const FINISHED = 'datastar-finished';
const ERROR = 'datastar-error';

/**
 * IPC Action Plugin for Datastar
 * Allows calling Wails backend methods and handling responses as Datastar events
 * 
 * @type {import('datastar').ActionPlugin}
 */
export const IPC = {
  type: 3,
  name: 'ipc',
  
  /**
   * Executes a Wails method and dispatches the resulting events to Datastar
   * @param {Object} ctx - Datastar action context
   * @param {string} method - The Wails App method to call
   * @param {...any} args - Arguments to pass to the Wails method
   * @returns {Promise<any>} - The result from the Wails method
   */
  async fn(ctx, method, ...args) {
    const elId = ctx.el.id;
    console.log('elid', elId);
    
    // Dispatch the started event
        console.log('starting');
    dispatchSSE(STARTED, elId, {});
    
    try {
      // Call the Wails method
      const raw = await window.go.main.App[method](...args);
      console.log("Raw response:", raw);
      
      // Wails automatically converts []byte to Base64-encoded string
      // Decode it and parse as JSON
      const decoded = typeof raw === 'string' ? atob(raw) : raw;
      const events = JSON.parse(decoded);
      console.log("events", events);
      
      // Dispatch each event to Datastar
      for (const { type, args } of events) {
        console.log('sse fragment', type, args);
        dispatchSSE(type, elId, args);
      }
      
      // Find signal data in the events if present
      const signalEvent = events.find(e => e.type === 'datastar-merge-signals');
      if (signalEvent && signalEvent.args && signalEvent.args.json) {
        // Return the parsed signal data as the action result
        console.log('sse signals')
        return JSON.parse(signalEvent.args.json);
      }
      
      return null;
    } catch (e) {
      // Handle errors
      console.error(`IPC action error (${method}):`, e);
      dispatchSSE(ERROR, elId, { message: e?.message ?? 'IPC error' });
      throw e;
    } finally {
      // Always dispatch the finished event
      dispatchSSE(FINISHED, elId, {});
    }
  },
};

// For Wails push events - register this in index.html
export function setupIpcPushEvents() {
  window.runtime.EventsOn("datastar-ipc", (_id, data) => {
    const events = JSON.parse(data);
    events.forEach(({type, args}) => {
      dispatchSSE(type, "__push__", args);
    });
  });
}
