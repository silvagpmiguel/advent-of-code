YEAR=$1
DAY=$2

if [ -z "$YEAR" ] || [ -z "$DAY" ]; then
    echo "Error: npm run day <YEAR> <DAY>"
else
    if ! cat ./libs/$YEAR/index.ts | grep -q "day$DAY"; then 
        echo "export * from './day$DAY'" >> ./libs/$YEAR/index.ts
    fi
    touch ../input/$YEAR/$DAY.1 ../input/$YEAR/$DAY.2
    echo "
import { StringArrayInputParts, Day } from '@advent/utils'

export const day${DAY}_${YEAR} = ({ part1, part2 }: StringArrayInputParts): Day => ({
  part1: (): number => 0,
  part2: (): number => 0,
})
" > ./libs/${YEAR}/day${DAY}.ts
fi