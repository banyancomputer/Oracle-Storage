use std::ffi::CStr;
use std::os::raw::c_char;
use std::io::Read;


pub struct ObaoData {
    pub obao: *const c_char,
    pub hash: *const c_char,
    pub obao_len: usize,
    pub hash_len: usize,
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
    let obao_len = obao.len();
    let hash_len = 32;

    // Reference the contents of obao and hash
    let obao_ref: &[u8] = &obao;
    // let hash_ref: &[u8; 32] = &hash.as_bytes();
    let mut hash_vec = Vec::with_capacity(hash_len);
    hash_vec.extend_from_slice(&hash.as_bytes()[..]);
    let hash_ref: &[u8] = &hash_vec;

    // print the contents of hash_ref to see if it is correct.
    Box::new(ObaoData {
        obao: obao_ref.as_ptr() as *const c_char,
        hash: hash_ref.as_ptr() as *const c_char,
        obao_len,
        hash_len
    })
}
