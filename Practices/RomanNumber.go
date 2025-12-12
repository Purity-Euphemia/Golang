use crate::RomanDigit::*;

#[derive(Copy, Clone, Debug, PartialEq, Eq)]
pub enum RomanDigit {
    Nulla,
    I,
    V,
    X,
    L,
    C,
    D,
    M,
}
pub struct RomanNumber(pub Vec<RomanDigit>);

impl From<u32> for RomanNumber {
    fn from(mut value: u32) -> Self {
        if value == 0 {
            return RomanNumber(vec![Nulla]);
        }

        let table: &[(u32, &[RomanDigit])] = &[
            (1000, &[M]),
            (900, &[C, M]),
            (500, &[D]),
            (400, &[C, D]),
            (100, &[C]),
            (90,  &[X, C]),
            (50,  &[L]),
            (40,  &[X, L]),
            (10,  &[X]),
            (9,   &[I, X]),
            (5,   &[V]),
            (4,   &[I, V]),
            (1,   &[I]),
        ];

        let mut digits = Vec::new();

        for &(val, repr) in table {
            while value >= val {
                value -= val;
                digits.extend_from_slice(repr);
            }
        }

        RomanNumber(digits)
    }
}
