use hyper::{Body, Method, Request};

use lazy_static::lazy_static;
use std::sync::atomic::{AtomicUsize, Ordering};
use std::sync::Arc;

#[tokio::main]
async fn main() {
    let http2 = std::env::var("HTTP2").map(|_| true).unwrap_or(false);

    let cnt = Arc::new(AtomicUsize::new(0));

    client(cnt.clone(), http2);

    loop {
        let prev_cnt = cnt.load(Ordering::Relaxed);
        tokio::time::sleep(std::time::Duration::from_secs(5)).await;
        let now_cnt = cnt.load(Ordering::Relaxed);
        println!("avg qps {}", (now_cnt - prev_cnt) / 5)
    }
}

lazy_static! {
    static ref URL: String = format!("http://localhost:{}", 1010);
}

fn client(cnt: Arc<AtomicUsize>, http2: bool) {
    let client = hyper::client::Client::builder()
        .http2_only(http2)
        .build_http();
    for _ in 0..16 {
        let client = client.clone();
        let cnt = cnt.clone();
        tokio::task::spawn(async move {
            loop {
                let req = Request::builder()
                    .method(Method::POST)
                    .uri(URL.clone())
                    .body(Body::from("Hello!"))
                    .unwrap();
                match client.request(req).await {
                    Ok(_) => (),
                    Err(err) => {
                        println!("err {:?}", err);
                        return;
                    }
                }
                cnt.fetch_add(1, Ordering::SeqCst);
            }
        });
    }
}
