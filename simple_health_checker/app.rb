require 'net/http'

host = (ENV['host'] || "echo-server-svc:1323")

while true
  $stdout.puts "start to request: http://#{host}/?from=health-checker"
  $stdout.flush
  uri = URI("http://#{host}/?from=health-checker")
  puts Net::HTTP.get(uri)
  sleep 1
end

