use battleship::*;
use serde::{Deserialize, Serialize};

#[derive(Serialize, Deserialize, Debug)]
pub struct Game {
    pub id: String,
    pub stake: i64,
    pub current_player_id: String,
    pub mediator: mediator::Mediator,
    pub winner_id: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct SettledGames {
    pub games: Vec<Game>,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct CreateGameRequest {
    pub player_name: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct JoinGameRequest {
    pub game_id: String,
    pub player_name: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct InitFieldRequest {
    pub game_id: String,
    pub field: [[structures::Status; structures::LEN as usize]; structures::LEN as usize],
}

#[derive(Serialize, Deserialize, Debug)]
pub struct MoveRequest {
    pub game_id: String,
    pub point: structures::Point,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct QuitGameRequest {
    pub game_id: String,
}

#[derive(Serialize, Deserialize, Debug)]
pub struct GetGameRequest {
    pub game_id: String,
}

impl Game {
    pub fn new(player_a_id: &String, player_a_name: &String, stake: i64) -> Game {
        let player_a = player::Player::new(&player_a_id, &player_a_name);
        let mediator = mediator::Mediator::new(player_a);
        Game {
            id: "0000000001".to_string(), // Wil later use nanoid here to give the game a random id
            stake,
            mediator,
            current_player_id: "".to_string(),
            winner_id: "".to_string(),
        }
    }

    pub fn join(&mut self, player_b_id: &String, player_b_name: &String) -> &Game {
        let player_b = player::Player::new(&player_b_id, &player_b_name);
        self.mediator.join(player_b);
        self.current_player_id = player_b_id.to_string();
        self
    }

    pub fn init(&mut self, player_id: String, field: structures::Field) {
        self.mediator.init(player_id, field);
    }

    pub fn is_ready(&self) -> bool {
        self.mediator.player_a.is_some()
            && self.mediator.player_b.is_some()
            && self.mediator.player_a.as_ref().unwrap().own_field.field
                != [[structures::Status::Empty; 12]; 12]
            && self.mediator.player_b.as_ref().unwrap().own_field.field
                != [[structures::Status::Empty; 12]; 12]
    }

    pub fn is_player(&self, caller_agent_id: String) -> bool {
        (self.mediator.player_a.is_some()
            && self.mediator.player_a.as_ref().unwrap().id == caller_agent_id)
            || (self.mediator.player_b.is_some()
                && self.mediator.player_b.as_ref().unwrap().id == caller_agent_id)
    }

    pub fn start(&mut self) {
        // TODO
    }

    pub fn player_move(&mut self, player_id: String, point: structures::Point) {
        // TODO
    }

    pub fn won(&mut self, player_id: &String) {
        self.winner_id = player_id.to_string();
    }

    pub fn is_settled(&self) -> bool {
        self.winner_id != "".to_string()
    }
}
