

pub fn vector_as_u8_array(vector: Vec<u8>) -> [u8;37] {
    let mut arr = [0u8;37];
    for (place, element) in arr.iter_mut().zip(vector.iter()) {
        *place = *element;
    }
    arr
}