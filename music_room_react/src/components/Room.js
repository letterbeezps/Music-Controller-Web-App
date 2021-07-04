/* eslint-disable no-unused-vars */
import React, { Component } from 'react';
import { Grid, Button, Typography } from "@material-ui/core";
import { Link } from "react-router-dom";
import CreateRoomPage from './CreateRoomPage';
import MusicPlayer from './MusicPlayer';

export default class Room extends Component {
    constructor(props) {
        super(props);
        this.state = {
            votesToSkip: 2,
            guestCanPause: false,
            isHost: false,
            showSettings: false,
            spotifyAuthenticated: false,
            song: {},
        };
        this.roomCode = this.props.match.params.roomCode;
        
        this.leaveButtonPressed = this.leaveButtonPressed.bind(this);
        this.updateShowSettings = this.updateShowSettings.bind(this);
        this.renderSettingsButton = this.renderSettingsButton.bind(this);
        this.renderSettings = this.renderSettings.bind(this);
        this.getRoomDetails = this.getRoomDetails.bind(this);
        this.authenticateSpotify = this.authenticateSpotify.bind(this);
        this.getCurrentSong = this.getCurrentSong.bind(this)

        this.getRoomDetails();
    }

    componentDidMount() {
        this.interval = setInterval(this.getCurrentSong, 1000)
    }

    componentWillUnmount() {
        clearInterval(this.interval)
    }

    getRoomDetails() {
        const requestOptions = {
            headers: {'Content-Type': 'application/json'},
            mode: 'cors',
            credentials: "include",
        };
        fetch("http://192.168.199.133:9898/api/v1/room?roomCode=" + this.roomCode, requestOptions).then((response) =>
            response.json()
        ).then((data) => {

            console.log(data)

            if (data.code !== 200) {
                console.log("没有本地session")
                this.props.leaveRoomCallback();
                this.props.history.push("/");
            } else {
                this.setState({
                    votesToSkip: data.data.room.votes_to_skip,
                    guestCanPause: data.data.room.guest_can_pause,
                    isHost: data.data.is_host,
                });
                if (this.state.isHost) {
                    this.authenticateSpotify();
                    }
                }
        });
    }

    authenticateSpotify() {
        const requestOptions = {
            headers: {'Content-Type': 'application/json'},
            mode: 'cors',
            credentials: "include",
        };
        fetch("http://192.168.199.133:9898/api/spotify/Is_authenticated", requestOptions).then((response) => 
            response.json()
        ).then((data) => {
            console.log(data);
            this.setState({
                spotifyAuthenticated: data.data.status
            });
            if (!data.data.status) {
                fetch("http://192.168.199.133:9898/api/spotify/Get_auth_url", requestOptions).then((response) => 
                    response.json()
                ).then((data) => {
                    window.location.replace(data.data.url);
                });
            }
        })
    }

    getCurrentSong() {
        const requestOptions = {
            headers: {'Content-Type': 'application/json'},
            mode: 'cors',
            credentials: "include",
        };
        fetch("http://192.168.199.133:9898/api/spotify/current_song", requestOptions).then((response) => 
            response.json()
        ).then((data) => {
            if (data.code === 200) {
                this.setState({
                    song: data.data,
                })
            } else {
                this.setState({
                    song: {},
                })
            }
        })
    }

    leaveButtonPressed() {
        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            mode: 'cors',
            credentials: "include",
        };
        fetch("http://192.168.199.133:9898/api/v1/Leave_room", requestOptions).then((_response) => {
            this.props.leaveRoomCallback();
            this.props.history.push('/');
        });
    }

    updateShowSettings(value) {
        this.setState({
            showSettings: value,
        });
    }

    renderSettingsButton() {
        return (
            <Grid item xs={12} align="center">
                <Button 
                    variant="contained"
                    color="primary"
                    onClick={() => this.updateShowSettings(true)}
                >
                    设置
                </Button>
            </Grid>
        );
    }

    renderSettings() {

        return (
            <Grid
                container spacing={1}
            >
                <Grid item xs={12} align="center">
                    <CreateRoomPage 
                        update={true}
                        votesToSkip={this.state.votesToSkip} 
                        guestCanPause={this.state.guestCanPause}
                        roomCode={this.roomCode}
                        updateCallback={this.getRoomDetails}
                    />
                </Grid>
                <Grid item xs={12} align="center">
                    <Button 
                        variant="contained"
                        color="secondary"
                        onClick={() => this.updateShowSettings(false)}
                    >
                        关闭
                    </Button> 
                </Grid> 
            </Grid>
        )

    }

    render() {
        if (this.state.showSettings) {
            return this.renderSettings();
        }
        return <Grid container spacing={1}>
            <Grid item xs={12} align="center">
                <Typography variant="h4" component="h4">
                    Code: {this.roomCode}
                </Typography>
            </Grid>
            <MusicPlayer {...this.state.song} />
            {/* <Grid item xs={12} align="center">
                <Typography variant="h6" component="h6">
                    Votes: {this.state.votesToSkip}
                </Typography>
            </Grid>
            <Grid item xs={12} align="center">
                <Typography variant="h6" component="h6">
                    Guest Can pause: {this.state.guestCanPause.toString()}
                </Typography>
            </Grid>
            <Grid item xs={12} align="center">
                <Typography variant="h6" component="h6">
                    Host: {this.state.isHost.toString()}
                </Typography>
            </Grid> */}
            
            {this.state.isHost ? this.renderSettingsButton() : null}
            <Grid item xs={12} align="center">
                <Button variant="contained" color="secondary" component={Link} onClick={this.leaveButtonPressed}>
                    回到首页
                </Button>
            </Grid>
        </Grid>
    }
}