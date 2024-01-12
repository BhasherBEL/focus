// Prevents additional console window on Windows in release, DO NOT REMOVE!!
#![cfg_attr(not(debug_assertions), windows_subsystem = "windows")]

fn main() {
  tauri::Builder::default()
    .invoke_handler(tauri::generate_handler![get_backend_url])
    .run(tauri::generate_context!())
    .expect("error while running tauri application");
}

#[tauri::command]
fn get_backend_url() -> String {
  std::env::var("PUBLIC_BACKEND_URL").unwrap_or_else(|_| "http://localhost:3000".to_string())
}