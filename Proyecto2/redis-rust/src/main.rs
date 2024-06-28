#[macro_use] extern crate rocket;

use rocket::serde::json::Json;
use rocket::serde::{Deserialize, Serialize};
use redis::Commands;

#[derive(Deserialize, Serialize)]
struct Data {
    key: String,
    value: String,
}

#[post("/set", format = "json", data = "<data>")]
async fn set_data(data: Json<Data>) -> Result<&'static str, &'static str> {
    // Crear cliente de redis
    let client = redis::Client::open("redis://34.122.189.228:6379/")
        .map_err(|_| "Failed to create Redis client")?;

    // Conexion a redis
    let mut con = client.get_connection()
        .map_err(|_| "Failed to connect to Redis")?;

    // Insertar hash en redis
    let _: () = con.hincr(&data.key, &data.value, 1)
        .map_err(|_| "Failed to set data in Redis")?;
    Ok("Data set")
}

#[launch]
fn rocket() -> _ {
    rocket::build().mount("/", routes![set_data])
}