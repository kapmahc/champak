require 'highline'
require 'csv'

namespace :magento do
  desc 'Import products from magento v1 cvs'
  task products: :environment do
    # http://docs.magento.com/m1/ce/user_guide/store-operations/data-export.html
    # https://www.atensoftware.com/p187.php
    Dir.glob("#{Rails.root}/tmp/magento/*.csv") do |fn|
      cli = HighLine.new
      unless cli.agree("load data from file #{fn}?  ")
        next
      end

      CSV.foreach(fn) do |row|
        puts "find product #{row}"

      end
    end
  end

end
