FROM gradle:8.13.0-jdk21-alpine AS build
COPY --chown=gradle:gradle . /home/gradle/src
WORKDIR /home/gradle/src
RUN gradle build --no-daemon 

FROM eclipse-temurin:21-jdk-alpine

EXPOSE 8080

RUN mkdir /app

COPY --from=build /home/gradle/src/app/build/libs/app.jar /app/app.jar

ENTRYPOINT ["java", "-jar","/app/app.jar"]
