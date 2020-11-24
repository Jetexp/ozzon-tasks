use std::io::prelude::*;
use std::fs::File;
use std::io::{BufRead, BufReader, Read};
use std::collections::HashMap;

fn read<R: Read>(io: R) -> (i32, HashMap<i32, i32>) {
    let mut hash: HashMap<i32, i32> = HashMap::new();
    let reader = BufReader::new(io);
    let mut target = 0;
    let mut first = true;

    for line in reader.lines() {
        if first {
            first = false;

            target = line
                .unwrap()
                .parse()
                .expect("Wanted a number");

            continue;
        }

        for word in line.unwrap().split_whitespace() {
            let num = word
                .parse()
                .expect("Wanted a number");

            if num > target {
                continue
            }

            hash.insert(num, 1 + if hash.contains_key(&num) { hash[&num] } else { 0 });
        }
    }

    (target, hash)
}

fn f() -> bool {
    let (target, hash) = read(File::open("input.txt").expect("Cannot open input.txt"));

    //println!("{}", target);
    //println!("{:?}", hash);

    for (current, count) in hash.iter() {
        //println!("{}, {}", k, v)

        let need = &target - current;

        if hash.contains_key(&need) {
            if need == *current {
                if count <= &1 {
                    continue
                }
            }

            return true;
        }
    }

    return false
}

// very fast
fn main() {
    let mut result = "0";

    if f() {
        result = "1"
    }

    let mut fs = File::create("output.txt").unwrap();
    fs.write(result.as_bytes()).unwrap();
}