#[allow(unused_imports)]
use chrono::{Local, TimeZone};
use getopts::Options;
use std::env;

const HINTOPTION: &str = "h";

fn print_usage(program: &str, opts: Options) {
    let brief = format!("Usage: {} FILE [options]", program);
    print!("{}", opts.usage(&brief));
}

fn main() {
    let args: Vec<String> = env::args().collect();

    let program = args[0].clone();
    let mut opts = Options::new();
    opts.optopt("l", "", "datetime layout", "DateLayout");
    opts.optflag(HINTOPTION, "help", "print this help menu");

    let matches = match opts.parse(&args[1..]) {
        Ok(m) => m,
        Err(f) => {
            panic!("{}", f)
        }
    };
    if matches.opt_present(HINTOPTION) {
        println!("hint !!!!!!!!!!11")
    }
    let lo = matches.opt_str("l");
    let layout = lo.unwrap_or("%Y-%m-%d %H:%M:%S".to_string());
    let input = if !matches.free.is_empty() {
        matches.free.join(" ")
    } else {
        print_usage(&program, opts);
        return;
    };
    match input.parse::<i64>() {
        Ok(ut) => {
            print!("{}", Local.timestamp(ut, 0).format(&layout))
        }
        Err(_f) => datetime_to_unix(&input),
    };
}

fn datetime_to_unix(dt: &str) {
    let layouts = vec!["%Y-%m-%d %H:%M:%S", "%Y%m%d %H:%M:%S", "%Y/%m/%d %H:%M:%S"];
    for l in layouts {
        match Local.datetime_from_str(dt, l) {
            Ok(d) => {
                print!("{}", d.timestamp());
                return;
            }
            Err(_f) => {}
        }
    }
}
