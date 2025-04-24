// Wails IPC integration for Datastar

// Global function to call Wails methods and update the DOM with results
window.callWailsMethod = async function(method, ...args) {
  console.log(`Calling Wails method: ${method}`, args);
  
  try {
    // Show loading state by dispatching a custom event
    document.dispatchEvent(new CustomEvent('datastar-loading-start'));
    
    // Call the Go backend method via Wails IPC
    const result = await window.go.main.App[method](...args);
    console.log(`Result from ${method}:`, result);
    
    // For Inc method, update the count display
    if (method === 'Inc') {
      const countEl = document.getElementById('count');
      if (countEl) {
        countEl.textContent = result;
      }
    }
    
    // Hide loading state
    document.dispatchEvent(new CustomEvent('datastar-loading-end'));
    
    return result;
  } catch (err) {
    console.error(`Error calling Wails method ${method}:`, err);
    
    // Hide loading state even on error
    document.dispatchEvent(new CustomEvent('datastar-loading-end'));
    
    return null;
  }
};