require 'sinatra'
require "digest/md5"

set :port, 5558

get '/' do
  "#{Digest::MD5.hexdigest(Time.now.to_i.to_s)}"
end