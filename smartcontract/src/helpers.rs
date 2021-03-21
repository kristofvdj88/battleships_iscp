use iota_sc_utils::{getter::Getter, params};

pub fn vector_as_u8_array(vector: Vec<u8>) -> [u8; 37] {
    let mut arr = [0u8; 37];
    for (place, element) in arr.iter_mut().zip(vector.iter()) {
        *place = *element;
    }
    arr
}

pub fn get_struct<TContext: Getter, T: serde::de::DeserializeOwned>(
    param_name: &str,
    ctx: &TContext,
) -> Result<T, serde_json::Error> {
    let json_str = params::get_string(param_name, ctx);
    serde_json::from_str(&json_str)
}
