export function smartFloat(anchorId: string, targetId: string) {
    const anchorEl = document.getElementById(anchorId);
    const targetEl = document.getElementById(targetId);
  
    if (!(anchorEl && targetEl)) return;
    const centerX = anchorEl.offsetLeft + anchorEl.clientWidth / 2;
    const centerY = anchorEl.offsetTop + anchorEl.clientHeight / 2;
  
    targetEl.style.left = `${centerX}px`;
    targetEl.style.top = `${centerY}px`;
  }
  