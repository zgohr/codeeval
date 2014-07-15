File.readlines(ARGV[0]).each do |l|
    n, l = l.strip!.split(',')
    strings = l.chars.sort.repeated_permutation(n.to_i).map(&:join).uniq
    puts strings.join(',')
end
