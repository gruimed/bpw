package otel;

import org.springframework.boot.Banner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class DiceApplication {
  public static void main(String[] args) {

    try {
      // The newInstance() call is a work around for some
      // broken Java implementations

      Class.forName("com.mysql.jdbc.Driver");
    } catch (Exception ex) {
        // handle the error
    }


    SpringApplication app = new SpringApplication(DiceApplication.class);
    app.setBannerMode(Banner.Mode.OFF);
    app.run(args);
  }
}
