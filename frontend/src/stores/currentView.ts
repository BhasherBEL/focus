import { writable } from "svelte/store";
import type { View } from "./interfaces";

export default writable(null as View | null);