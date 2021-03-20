use crate::field;
use crate::structures;
use field::GameField;
use structures::{Move, Point, Status, Field};
use serde::{Serialize, Deserialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Player {
    pub id: String,
    pub name: String,
    pub own_field: GameField,
    pub enemy_field: GameField,    
}

impl Player {
    pub fn new(id: &String, name: &String) -> Player {
        Player {
            id: id.to_string(),
            name: name.to_string(),
            own_field:GameField::new(),
            enemy_field:GameField::new(),            
        }
    }

    pub fn init(&mut self, field: Field) {
        self.own_field.field = field;
    }

    pub fn random_init(&mut self) {
        self.own_field.generate_random_field();
    }

    pub fn enemy_attack(&mut self, point: Point) -> Move {
        match self.own_field.get_cell_value(point) {
            Status::Ship => {
                self.own_field.draw_cell(point, Status::Kill);
                self.own_field.sink_ship();
                Move::Kill(point)
            }
            _ => Move::Miss(point),
        }
    }

    pub fn player_move(&mut self, result: &Move) {
        match result {
            Move::Kill(point) => {
                println!("HIT!\n");
                self.enemy_field.draw_cell(*point, Status::Kill);
            }
            Move::Miss(point) => {
                println!("MISS!\n");
                self.enemy_field.draw_cell(*point, Status::Bound);
            }
        }
    }
}
