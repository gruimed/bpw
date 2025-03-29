package otel;

import java.util.Optional;
import java.util.concurrent.ThreadLocalRandom;
import java.util.Arrays;
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
//import java.net.URI;
import java.io.IOException;
import java.lang.InterruptedException;

@RestController
public class RollController {
  private static final Logger logger = LoggerFactory.getLogger(RollController.class);

  @GetMapping("/rolldice")
  public String index(@RequestParam("player") Optional<String> player) {
    /* 
    int[] arr = new int[1000000];
    for (int i = 0 ; i < 1000000 ; i++) {
      arr[i] = 1000000 - i;
    }

    Arrays.sort(arr);

    */

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
 


    int result = this.getRandomNumber(1, 6);
    if (player.isPresent()) {
//      logger.info("{} is rolling the dice: {}", player.get(), result);
    } else {
//      logger.info("Anonymous player is rolling the dice: {}", result);
    }
    return Integer.toString(result);
  }

  public int getRandomNumber(int min, int max) {
    return ThreadLocalRandom.current().nextInt(min, max + 1);
  }
}
