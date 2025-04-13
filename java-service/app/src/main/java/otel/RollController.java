package otel;

import java.util.Optional;
import java.util.concurrent.ThreadLocalRandom;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestParam;
import org.springframework.web.bind.annotation.RestController;

import java.net.*;
import java.net.http.*;

//import java.net.http.HttpRequest;
//import java.net.http.HttpClient;
//import java.net.http.HttpResponse;
import java.net.http.HttpClient.Version;
import java.net.http.HttpResponse.BodyHandlers;
import java.sql.Connection;
import java.sql.DriverManager;
import java.sql.ResultSet;
import java.sql.SQLException;
import java.sql.Statement;
//import java.net.URI;
import java.io.IOException;
import java.lang.InterruptedException;

import io.opentelemetry.api.GlobalOpenTelemetry;
import io.opentelemetry.api.trace.Tracer;
import io.opentelemetry.api.trace.Span;

@RestController
public class RollController {
private static final Logger logger = LoggerFactory.getLogger(RollController.class);
private static final Tracer tracer = GlobalOpenTelemetry.get().getTracerProvider().tracerBuilder("java-service-random").build();

  @GetMapping("/rolldice")
  public String index(@RequestParam("rolls") Optional<Integer> rolls, @RequestParam("load") Optional<String> load) {

    Integer rollsnumber = 1;
    if (rolls.isPresent()) {
      rollsnumber = rolls.get();
    }

    List<Integer> result = new ArrayList<Integer>(); 
    for (Integer i=0 ; i < rollsnumber ; i++) {
      result.add(rollonce(load));
    }

    return result.toString();
  }

  public Integer getRandomNumber(int min, int max) {

//    var span = tracer.spanBuilder("getRandom").startSpan();

    Integer result = ThreadLocalRandom.current().nextInt(min, max + 1);

//    span.end();
    return result;
  }

  private Integer rollonce(Optional<String> load) {

    if (load.isPresent() && load.get().indexOf('C') > -1) {
      int[] arr = new int[1000000];
      for (int i = 0 ; i < 1000000 ; i++) {
        arr[i] = 1000000 - i;
      }
  
      Arrays.sort(arr);
  
    }


    if (load.isPresent() && load.get().indexOf('E') > -1) {
      try {
        HttpRequest request = HttpRequest.newBuilder(new URI("http://echo-service:8088/payload?io_msec=10")).build();
        HttpClient client = HttpClient.newBuilder()
        .version(Version.HTTP_1_1)
        .build();

        HttpResponse<String> response = client.send(request, BodyHandlers.ofString());
      } catch (URISyntaxException e) {

      } catch (IOException e) {
        
      } catch (InterruptedException e) {
      }
      
    }

    if (load.isPresent() && load.get().indexOf('D') > -1) {
      Statement stmt = null;
      ResultSet rs = null;

      try {
        Connection conn = DriverManager.getConnection("jdbc:mysql://pinba/?user=root&password=");
        stmt = conn.createStatement();
        rs = stmt.executeQuery("SELECT now()");
      } catch (SQLException ex) {
        logger.error("DB error: " + ex.getMessage());
        return 0;
      }
      finally {
        if (rs != null) {
          try {
              rs.close();
          } catch (SQLException sqlEx) { } // ignore
  
          rs = null;
        }
  
        if (stmt != null) {
          try {
              stmt.close();
          } catch (SQLException sqlEx) { } // ignore
  
          stmt = null;
        }
      }
    }

    return this.getRandomNumber(1, 6);

  }
}
