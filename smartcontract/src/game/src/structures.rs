use std::collections::HashMap;
use serde::{Serialize, Deserialize};

#[derive(Debug)]
pub enum ShipDirection {
    Horizontal,
    Vertical,
}

#[derive(Debug)]
pub enum Move {
    Miss(Point),
    Kill(Point),
}

#[derive(Debug)]
pub enum Direction {
    Up,
    Right,
    Down,
    Left,
}

#[derive(Debug, Copy, Clone)]
pub struct Ship {
    pub size: u8,
    pub direction: &'static ShipDirection,
    pub start_point: Point,
}

impl Ship {
    pub fn get_all() -> HashMap<u8, u8> {
        /// @todo replace HashMap with Array
        /// Regarding to https://users.rust-lang.org/t/rust-battleship-console-game/31028/2
        let mut ships = HashMap::new();
        let keys: [u8; 4] = [1, 2, 3, 4];
        let mut values = keys.iter().rev();

        for &key in keys.iter() {
            ships.insert(key, *values.next().unwrap());
        }
        ships
    }
}

#[derive(Serialize, Deserialize, Debug, Copy, Clone, PartialEq)]
pub enum Status {
    Empty,
    Ship,
    Bound,
    Kill,
}

pub struct Draw {
    pub start_point: Point,
    pub path: Vec<(Direction, u8)>,
    pub draw_status: Status,
    pub allowed_status: Vec<Status>,
}

#[derive(Serialize, Deserialize, Debug, Clone, Copy)]
pub struct Point {
    pub row: u8,
    pub column: u8,
}
impl PartialEq for Point {
    fn eq(&self, other: &Point) -> bool {
        self.row == other.row && self.column == other.column
    }
}

impl Point {
    pub fn go_to(&mut self, direction: &Direction) -> &mut Self {
        match direction {
            Direction::Up => self.row -= 1,
            Direction::Left => self.column -= 1,
            Direction::Right => self.column += 1,
            Direction::Down => self.row += 1,
        }
        self
    }
}

pub const LEN: u8 = 12;

pub type Field = [[Status; LEN as usize]; LEN as usize];
