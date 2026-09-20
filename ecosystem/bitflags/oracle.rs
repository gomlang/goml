use bitflags::{bitflags, parser, Flags};
use std::fmt::Write;
use std::io::{self, BufRead};

bitflags! {
    #[derive(Clone, Copy)]
    struct Access: u8 {
        const NONE = 0;
        const READ = 1;
        const WRITE = 2;
        const BOTH = 3;
        const ALIAS = 1;
        const GROUP = 12;
        const _ = 64;
    }
    #[derive(Clone, Copy)]
    struct Overlap: u8 { const A = 1; const AB = 3; }
    #[derive(Clone, Copy)]
    struct Empty: u8 {}
    #[derive(Clone, Copy)]
    struct External: u64 { const A = 1; const _ = !0; }
    #[derive(Clone, Copy)]
    struct Wide: u64 { const LOW = 1; const HIGH = 1 << 63; const BOTH = 1 | (1 << 63); }
    #[derive(Clone, Copy)]
    struct Medium: u16 { const LOW = 1; const HIGH = 1 << 15; const GROUP = 240; }
    #[derive(Clone, Copy)]
    struct Word: u32 { const LOW = 1; const HIGH = 1 << 31; const GROUP = 240; }
}

fn parsed<F: Flags>(value: Result<F, parser::ParseError>) -> String
where
    F::Bits: Into<u64>,
{
    value
        .map(|v| v.bits().into().to_string())
        .unwrap_or_else(|_| "error".into())
}

fn run<F>(a: F, b: F, input: &str) -> String
where
    F: Flags + Copy,
    F::Bits: Into<u64> + parser::WriteHex + parser::ParseHex,
{
    let mut text = String::new();
    let mut truncated = String::new();
    let mut strict = String::new();
    parser::to_writer(&a, &mut text).unwrap();
    parser::to_writer_truncate(&a, &mut truncated).unwrap();
    parser::to_writer_strict(&a, &mut strict).unwrap();
    let mut fields = vec![
        a.bits().into().to_string(),
        F::all().bits().into().to_string(),
        F::all_named().bits().into().to_string(),
        a.known_bits().into().to_string(),
        a.unknown_bits().into().to_string(),
        a.is_empty().to_string(),
        a.is_all().to_string(),
        F::from_bits(a.bits()).is_some().to_string(),
        a.union(b).bits().into().to_string(),
        a.intersection(b).bits().into().to_string(),
        a.difference(b).bits().into().to_string(),
        a.symmetric_difference(b).bits().into().to_string(),
        a.complement().bits().into().to_string(),
        a.contains(b).to_string(),
        a.intersects(b).to_string(),
        text,
        truncated,
        strict,
    ];
    let mut names = String::new();
    for (name, _) in a.iter_names() {
        write!(names, "{name},").unwrap();
    }
    fields.push(names);
    let mut values = String::new();
    for value in a.iter() {
        write!(values, "{},", value.bits().into()).unwrap();
    }
    fields.push(values);
    let mut equal = String::new();
    for name in a.iter_equal_names() {
        write!(equal, "{name},").unwrap();
    }
    fields.push(equal);
    fields.push(parsed(parser::from_str::<F>(input)));
    fields.push(parsed(parser::from_str_truncate::<F>(input)));
    fields.push(parsed(parser::from_str_strict::<F>(input)));
    fields.join("\t")
}

fn main() {
    for line in io::stdin().lock().lines() {
        let line = line.unwrap();
        let fields: Vec<_> = line.split('\t').collect();
        assert_eq!(fields.len(), 4);
        let a = u64::from_str_radix(fields[1], 16).unwrap();
        let b = u64::from_str_radix(fields[2], 16).unwrap();
        let bytes: Vec<u8> = (0..fields[3].len())
            .step_by(2)
            .map(|i| u8::from_str_radix(&fields[3][i..i + 2], 16).unwrap())
            .collect();
        let text = String::from_utf8(bytes).unwrap();
        let result = match fields[0] {
            "access" => run(
                Access::from_bits_retain(a as u8),
                Access::from_bits_retain(b as u8),
                &text,
            ),
            "overlap" => run(
                Overlap::from_bits_retain(a as u8),
                Overlap::from_bits_retain(b as u8),
                &text,
            ),
            "empty" => run(
                Empty::from_bits_retain(a as u8),
                Empty::from_bits_retain(b as u8),
                &text,
            ),
            "external" => run(
                External::from_bits_retain(a),
                External::from_bits_retain(b),
                &text,
            ),
            "wide" => run(Wide::from_bits_retain(a), Wide::from_bits_retain(b), &text),
            "medium" => run(
                Medium::from_bits_retain(a as u16),
                Medium::from_bits_retain(b as u16),
                &text,
            ),
            "word" => run(
                Word::from_bits_retain(a as u32),
                Word::from_bits_retain(b as u32),
                &text,
            ),
            _ => panic!("unknown profile"),
        };
        println!("{result}");
    }
}
