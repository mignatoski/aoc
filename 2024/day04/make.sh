run ()
{
         python3 main.py
}

dbuild ()
{
        tag=${PWD##*/}
        sudo docker build -t aoc-$tag .
}

drun () {
        tag=${PWD##*/}
        sudo docker run --rm aoc-$tag
}

"$@"
