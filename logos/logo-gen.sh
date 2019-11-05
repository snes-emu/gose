sizes=('16' '32' '48' '64' '128' '256' '512' '1024')

for size in ${sizes[*]}
do
convert -size "${size}x${size}" -background none -extent "${size}x${size}" -gravity Center logo.svg "logo${size}.png"
done