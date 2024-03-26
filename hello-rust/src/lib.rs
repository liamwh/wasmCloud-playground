wit_bindgen::generate!({
    world: "hello",
    exports: {
        "wasi:http/incoming-handler": HttpServer,
    },
});

use exports::wasi::http::incoming_handler::Guest;
use wasi::http::types::*;

struct HttpServer;

impl Guest for HttpServer {
    fn handle(_request: IncomingRequest, response_out: ResponseOutparam) {
        let response = OutgoingResponse::new(Fields::new());
        response
            .set_status_code(200)
            .expect("setting status code to 200 should never fail");
        let response_body = response.body().unwrap();
        response_body
            .write()
            .expect("the first time we call write should always succeed")
            .blocking_write_and_flush(b"Hello from Rust!\n")
            .expect("a byte string literal smaller than 4096 bytes should always succeed");
        OutgoingBody::finish(response_body, None).expect("failed to finish response body");
        ResponseOutparam::set(response_out, Ok(response));
    }
}
