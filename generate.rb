require 'pp'
require 'date'

years = (2000...2084).map { |year| Time.gm(year).to_i }
dst = (2000...2084).map { |year| Date.leap?(year) }

months = (1...13).map { |month| Time.gm(1970, month).to_i }
dst_months = (1...13).map { |month| Time.gm(2004, month).to_i - Time.gm(2004).to_i }

pp dst_months
