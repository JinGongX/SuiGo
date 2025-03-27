import React from "react";
import { ActionManager } from "../actions/manager";
import { ExcalidrawElement, PointerType } from "../element/types";
import { AppState, Zoom } from "../types";
export declare const SelectedShapeActions: ({ appState, elements, renderAction, elementType, }: {
    appState: AppState;
    elements: readonly ExcalidrawElement[];
    renderAction: ActionManager["renderAction"];
    elementType: ExcalidrawElement["type"];
}) => JSX.Element;
export declare const ShapesSwitcher: ({ canvas, elementType, setAppState, onImageAction, }: {
    canvas: HTMLCanvasElement | null;
    elementType: ExcalidrawElement["type"];
    setAppState: React.Component<any, AppState>["setState"];
    onImageAction: (data: {
        pointerType: PointerType | null;
    }) => void;
}) => JSX.Element;
export declare const ZoomActions: ({ renderAction, zoom, }: {
    renderAction: ActionManager["renderAction"];
    zoom: Zoom;
}) => JSX.Element;
