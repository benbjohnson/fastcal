require 'pp'
require 'date'

years = (1970...2084).map { |year| Time.gm(year).to_i }
dst = (1970...2084).map { |year| Date.leap?(year) }

months = (1...13).map { |month| Time.gm(1970, month).to_i }
dst_months = (1...13).map { |month| Time.gm(2004, month).to_i - Time.gm(2004).to_i }
days = (0...31).map { |day| day * 24 * 60 * 60 }
hours = (0...23).map { |hour| hour * 60 * 60 }

pp dst
