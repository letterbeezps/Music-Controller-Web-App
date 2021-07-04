/* eslint-disable no-unused-vars */
import React, { Component } from "react";
import { 
    BrowserRouter as Router, 
    Switch, 
    Route, 
    Link, 
    Redirect 
} from "react-router-dom"; 

import { Grid, Button, ButtonGroup, Typography } from "@material-ui/core";

import RoomJoinPage from "./RoomJoinPage";
import CreateRoomPage from "./CreateRoomPage";
import Room from "./Room";

export default class HomePage extends Component {
    // eslint-disable-next-line no-useless-constructor
    constructor(props) {
        super(props);
        this.state = {
            roomCode: null,
        }
        this.clearRoomCode = this.clearRoomCode.bind(this);
    }

    async componentDidMount() {
        const requestOptions = {
            headers: {'Content-Type': 'application/json'},
            mode: 'cors',
            credentials: "include",
        };
        fetch("http://192.168.199.133:9898/api/v1/User_in_room", requestOptions)
        .then((response) => response.json())
        .then((data) => {
            if (data.data.code !== "") {
                this.setState({
                    roomCode: data.data.code,
                });
            } 
        });
    }

    renderHomePage() {
        return (
            <Grid container spacing={3}>
                <Grid item xs={12} align="center">
                    <Typography variant="h3" compact="h3">
                        家庭聚会
                    </Typography>
                </Grid>
                <Grid item xs={12} align="center">
                    <ButtonGroup
                        disableElevation variant="contained"
                        color="primary"
                    >
                        <Button color="primary" to="/join" component={ Link }>
                            加入歌房
                        </Button>
                        <Button color="secondary" to="/create" component={ Link }>
                            创建歌房
                        </Button>
                    </ButtonGroup>
                </Grid>
            </Grid>
        );
    }

    clearRoomCode() {
        this.setState({
            roomCode: null,
        })
    }

    render() {
        return (
            <Router>
                <Switch>
                    <Route exact path="/" render={() => {
                        return this.state.roomCode ? (
                        <Redirect to={`/room/${this.state.roomCode}`} />
                        ) : (
                            this.renderHomePage()
                            );
                        }}
                    />
                    
                    <Route path="/join" component={RoomJoinPage} />

                    <Route path="/create" component={CreateRoomPage} />

                    <Route 
                        path="/room/:roomCode"
                        render={(props) => {
                            return <Room {...props} leaveRoomCallback={this.clearRoomCode} />
                        }}
                    />
                </Switch>
            </Router>
        )
    }
}