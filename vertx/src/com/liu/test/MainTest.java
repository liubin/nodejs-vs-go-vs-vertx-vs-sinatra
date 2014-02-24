package com.liu.test;

import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

import org.vertx.java.core.Handler;
import org.vertx.java.core.Vertx;
import org.vertx.java.core.VertxFactory;
import org.vertx.java.core.http.HttpServerRequest;

public class MainTest implements Runnable {

    public static void main(String[] args) {
        ExecutorService threadsPool = Executors.newFixedThreadPool(1);
        threadsPool.execute(new MainTest());
    }

    private class MyHandler implements Handler<HttpServerRequest> {

        @Override
        public void handle(HttpServerRequest req) {
            String str = MD5Utils
                    .md5(String.valueOf(System.currentTimeMillis()));
            req.response().headers()
                    .set("Content-Type", "text/html; charset=UTF-8");
            req.response().end(str);
        }
    }

    @Override
    public void run() {
        Vertx vertx = VertxFactory.newVertx();
        vertx.createHttpServer().requestHandler(new MyHandler()).listen(5556);
        System.out.println("Server running at http://127.0.0.1:5556/");
    }
}
