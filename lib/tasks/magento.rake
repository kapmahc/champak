require 'sequel'
require 'highline'

# https://magento2.atlassian.net/wiki/display/m1wiki/Catalog+Database+Tables
namespace :magento do
  desc 'export entries from from magento database'
  task :export do
    cli = HighLine.new
    puts "Connect to magento's database"
    host = cli.ask('mysql host? ') { |q| q.default = 'localhost' }
    port = cli.ask('port? ', Integer) { |q| q.default = 3306 }
    user = cli.ask('username? ') {|q| q.default = 'root'}
    name = cli.ask('database name? ') { |q| q.default = 'test' }
    password = cli.ask('password? ') do |q|
      q.default = '1q2w#E4r'
      q.echo = 'x'
    end


    Sequel.connect(adapter: :mysql2, host: host, port: port, user: user, password: password, database: name, encoding: 'utf8') do |db|

    end

  end

end
