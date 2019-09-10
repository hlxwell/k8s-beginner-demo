# === Base Image ===
FROM ruby:alpine
ENV LC_ALL C.UTF-8
# Specify timezone
ENV TZ Asia/Tokyo
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

# Home dir
ENV APP_HOME /app
RUN mkdir -p $APP_HOME
ENV RACK_ENV production

# COPY FILES
ADD ./app.rb $APP_HOME

# Change config
WORKDIR $APP_HOME
CMD ["ruby", "./app.rb"]
