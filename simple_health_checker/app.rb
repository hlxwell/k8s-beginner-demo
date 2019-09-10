require 'net/http'

while true
  puts "start to request: http://echo-server-svc:1323/?from=health-checker"
  uri = URI('http://echo-server-svc:1323/?from=health-checker')
  puts Net::HTTP.get(uri)
  sleep 1
end
