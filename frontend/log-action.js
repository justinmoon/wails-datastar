// log-action.js  – registers @log()
import { load, apply, PluginType } from
  'https://cdn.jsdelivr.net/gh/starfederation/datastar@v1.0.0-beta.11/bundles/datastar.js';

/** @type {import('datastar').ActionPlugin} */
const LogAction = {
  type: PluginType.Action,
  name: 'log',
  fn: (_ctx, method, ...args) => {
    console.log(`[DS-log] ${method}`, ...args);
  },
};

// Register and immediately apply to the current DOM
load(LogAction);
apply();
