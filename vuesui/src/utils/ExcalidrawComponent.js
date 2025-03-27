import React from 'react';
import { createRoot } from 'react-dom/client';
import Excalidraw from '@excalidraw/excalidraw';
//import '@excalidraw/excalidraw/dist/excalidraw.min.css';

class ExcalidrawComponent extends React.Component {
  render() {
    return React.createElement(Excalidraw, null);
  }
}

export default ExcalidrawComponent;

export function renderExcalidraw(container) {
  const root = createRoot(container);
  root.render(React.createElement(ExcalidrawComponent));
  return root;
}

export function unmountExcalidraw(root) {
  if (root) {
    root.unmount();
  }
}
