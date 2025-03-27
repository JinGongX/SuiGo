import { ExcalidrawElement } from "../element/types";
import { AppState, BinaryFiles, LibraryItems } from "../types";
import { ImportedDataState } from "./types";
import Library from "./library";
export declare const serializeAsJSON: (elements: readonly ExcalidrawElement[], appState: Partial<AppState>, files: BinaryFiles, type: "local" | "database") => string;
export declare const saveAsJSON: (elements: readonly ExcalidrawElement[], appState: AppState, files: BinaryFiles) => Promise<{
    fileHandle: import("browser-fs-access").FileSystemHandle | null;
}>;
export declare const loadFromJSON: (localAppState: AppState, localElements: readonly ExcalidrawElement[] | null) => Promise<import("./restore").RestoredDataState>;
export declare const isValidExcalidrawData: (data?: {
    type?: any;
    elements?: any;
    appState?: any;
} | undefined) => data is ImportedDataState;
export declare const isValidLibrary: (json: any) => any;
export declare const saveLibraryAsJSON: (libraryItems: LibraryItems) => Promise<void>;
export declare const importLibraryFromJSON: (library: Library) => Promise<void>;
