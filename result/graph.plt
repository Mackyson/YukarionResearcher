set grid
set terminal png font "MigMix 2M,7" size 1500,480
set datafile separator ","
unset key

set xtics rotate by -90

set output "meishi.png"
plot "data/meishi.csv" using 0:2:xtic(1) with boxes notitle

set terminal png font "MigMix 2M,7" size 1200,480

set output "meishi2.png"
plot "data/meishi2.csv" using 0:2:xtic(1) with boxes notitle

set output "doushi.png"
plot "data/doushi.csv" using 0:2:xtic(1) with boxes notitle

set terminal png font "MigMix 2M,8" size 640,480

set output "kandoushi.png"
plot "data/kandoushi.csv" using 0:2:xtic(1) with boxes notitle
set output "jodoushi.png"
plot "data/jodoushi.csv" using 0:2:xtic(1) with boxes notitle

set yrange [0:10]
set output "fukushi.png"
plot "data/fukushi.csv" using 0:2:xtic(1) with boxes notitle
set output "kandoushi.png"
plot "data/kandoushi.csv" using 0:2:xtic(1) with boxes notitle
unset output
