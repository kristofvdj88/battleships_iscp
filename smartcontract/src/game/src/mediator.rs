use crate::player;
use crate::structures;
use player::Player;
use std::io;
use structures::{Move, Point, ShipDirection, Field};
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Mediator {
    pub player_a: Option<Player>,
    pub player_b: Option<Player>,
}

fn read_line() -> String {
    ///read_line allocates a new String every time, while okay for user input,
    /// this could probably be better if we took a &mut String instead.
    let mut input_text = String::new();
    io::stdin()
        .read_line(&mut input_text)
        .expect("failed to read from stdin");
    input_text
}

fn parse_number(direction: &str) -> u8 {
    println!("Enter {}", direction);
    let line = read_line();
    let trimmed = line.trim();
    let mut number = 0;
    if let Ok(i) = trimmed.parse::<u8>() {
        number = i
    };
    number
}

fn get_point() -> Point {
    let row = parse_number("row");
    let column = parse_number("column");
    println!("User moved, row: {:?}, column: {:?}", row, column);
    Point { row, column }
}

impl Mediator {
    pub fn new(player_a: Player) -> Mediator {
        Mediator { player_a: Some(player_a), player_b: None }
    }

    pub fn join(&mut self, player_b: Player) {
        self.player_b = Some(player_b);
    }

    pub fn init(&mut self, player_id: String, field: Field) {
        if self.player_a.is_some() && self.player_a.as_ref().unwrap().id == player_id {
            self.player_a.as_mut().unwrap().init(field);
        } else if self.player_b.is_some() && self.player_b.as_ref().unwrap().id == player_id {
            self.player_b.as_mut().unwrap().init(field);
        }
    }

    pub fn human_move(&mut self) {
        let mut missed = false;
        while !missed {
            let point = get_point();
            let result = self.player_b.as_mut().unwrap().enemy_attack(point);
            self.player_a.as_mut().unwrap().player_move(&result);
            if let Move::Miss(_) = result {
                missed = true
            }
        }
    }

    pub fn ai_move(&mut self) {
        let mut missed = false;
        while !missed {
            let random_point = self
                .player_b
                .as_mut()
                .unwrap()
                .enemy_field
                .generate_random_point(&ShipDirection::Horizontal, 1);
            println!("AI moved, row: {:?}, column: {:?}", random_point.row, random_point.column);
            let result = self.player_a.as_mut().unwrap().enemy_attack(random_point);
            self.player_b.as_mut().unwrap().player_move(&result);
            if let Move::Miss(_) = result {
                missed = true
            }
        }
    }
}
