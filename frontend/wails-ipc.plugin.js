/**
 * Datastar "ipc" action plugin – call a Wails Go method and merge the
 * HTML it returns into the DOM.
 *
 * Usage in markup:  data-ipc="IncHTML"
 * Optional mods:
 *   - selector="CSS"   (defaults to "[id=count]")
 *   - mode="outer|inner|before|after|append|prepend" (default outer)
 */

import { load } from 'datastar';

// Define our plugin
const IPC = {
  type: 3,  // Plugin type 3 is for action plugins
  name: 'ipc',     // Plugin name (used in data-ipc attribute)

  /**  
   * Plugin handler function
   * ctx.el is the element that carries the data-ipc attribute
   */
  fn: async (ctx, methodName, { selector = '#count', mode = 'outer' } = {}) => {
    console.log('IPC plugin called with:', { methodName, selector, mode });
    const { el } = ctx;
    const elId = el.id;
    
    try {
      // Set loading state
      document.dispatchEvent(new CustomEvent('datastar-loading-start'));
      console.log('Loading state started');
      
      // Call the Wails Go method
      console.log('Calling Go method:', methodName);
      const html = await window.go.main.App[methodName]();
      console.log('Received HTML response:', html);
      
      // Find target element(s)
      const targets = selector
        ? document.querySelectorAll(selector)
        : [el];
      
      console.log('Found targets:', targets.length);
      
      // Update each target
      targets.forEach(target => {
        console.log('Updating target with mode:', mode);
        switch (mode) {
          case 'inner':
            target.innerHTML = html;
            break;
          case 'before':
            target.insertAdjacentHTML('beforebegin', html);
            break;
          case 'after':
            target.insertAdjacentHTML('afterend', html);
            break;
          case 'append':
            target.insertAdjacentHTML('beforeend', html);
            break;
          case 'prepend':
            target.insertAdjacentHTML('afterbegin', html);
            break;
          default: // 'outer'
            target.outerHTML = html;
            break;
        }
      });
      
    } catch (err) {
      console.error('Error in IPC plugin:', err);
      throw err;
    } finally {
      // Clear loading state
      document.dispatchEvent(new CustomEvent('datastar-loading-end'));
      console.log('Loading state ended');
    }
  }
};

// Register our plugin with Datastar
console.log('Registering IPC plugin with Datastar');
load(IPC);
console.log('IPC plugin registered');

// Add a direct click handler for debugging
document.addEventListener('click', (e) => {
  if (e.target.hasAttribute('data-ipc')) {
    console.log('Debug: Button with data-ipc clicked:', e.target);
    console.log('Debug: data-ipc value:', e.target.getAttribute('data-ipc'));
  }
});