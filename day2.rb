def safe(nums)
  deltas = nums.drop(1).zip(nums).map { |a, b| b - a }
  deltas.all? { |d| d > 0 && d < 4 } || deltas.all? { |d| d < 0 && d > -4 }
end

result1 = 0
result2 = 0
File.foreach('inputs/2.txt') do |line|
  nums = line.split(' ').map(&:to_i)

  if safe(nums)
    result1 += 1
  end

  (0..nums.length).each { |n|
    modified_nums = nums.take(n) + nums.drop(n + 1)
    if safe(modified_nums)
      result2 += 1
      break
    end
  }

end

puts result1
puts result2
