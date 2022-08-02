use std::ffi::CStr;
use std::os::raw::c_char;
use std::io::Read;


pub struct ObaoData {
    pub obao_data: *const c_char,
    pub hash_data: *const c_char,
    pub obao_data_len: usize,
    pub hash_data_len: usize,
}

/// Process a file into an ObaoData instance and return it.
#[no_mangle]
pub extern "C" fn obao_data(c_filepath: *const libc::c_char) -> Box<ObaoData> {
    // Extract the filepath from the C string.
    let filepath = unsafe { CStr::from_ptr(c_filepath).to_str().unwrap() };
    // Read the contents of the file as a vector of bytes.
    let mut file = std::fs::File::open(filepath).unwrap();
    let mut contents = Vec::new();
    file.read_to_end(&mut contents).unwrap();
    let (obao, hash) = bao::encode::outboard(&contents);
    let obao_data_len = obao.len();
    let hash_data_len = 32;
    // Reference the contents of obao and hash
    let obao_ref: &[u8] = &obao;
    let hash_ref: &[u8; 32] = &hash.as_bytes();
    // Get a pointer to the contents of obao and hash.
    let obao_data = obao_ref.as_ptr() as *const c_char;
    let hash_data = hash_ref.as_ptr() as *const c_char;
    // Create a new ObaoData instance.
    Box::new(ObaoData {
        obao_data,
        hash_data,
        obao_data_len,
        hash_data_len
    })
}

// // This is present so it's easy to test that the code works natively in Rust via `cargo test
// #[cfg(test)]
// pub mod test {
//     use super::*;
//
//     // This is meant to do the same stuff as the main function in the .go files.
//     #[test]
//     fn simulated_main_function () {
//         let filepath = CString::new("../../test/ethereum.pdf").unwrap();
//         let obao_data = obao_data(filepath.as_ptr());
//         // Read the hash as a string. Encode it in Hex.
//         let hash_str = unsafe { CStr::from_ptr(obao_data.hash_data).to_str().unwrap() };
//         let hash_hex = hex::encode(hash_str);
//         let obao_data_len = obao_data.obao_data_len;
//         println!("obao_data_len: {}", obao_data_len);
//         println!("hash_hex: {}", hash_hex);
//     }
//
// }
