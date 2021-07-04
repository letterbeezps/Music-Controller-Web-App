/* eslint-disable no-unused-vars */
import React, { Component } from "react";
import { TextField, Button, Grid, Typography, Grow } from "@material-ui/core";
import { Link } from "react-router-dom";

export default class RoomJoinPage extends Component {
    // eslint-disable-next-line no-useless-constructor
    constructor(props) {
        super(props);
        this.state = {
            roomCode: "",
            error: ""
        };

        this._handleTextFieldChange = this._handleTextFieldChange.bind(this);
        this._roomButtonPressed = this._roomButtonPressed.bind(this);
    }

    render() {
        return <Grid container spacing={1}>
            <Grid item xs={12} aligin="center">
                <Typography variant="h4" component="h4">
                    加入歌房
                </Typography>
            </Grid>
            <Grid item xs={12} aligin="center">
                <TextField
                    error="error"
                    label="Code"
                    placeholder="输入房间号"
                    value={this.state.roomCode}
                    helperText={this.state.error}
                    variant="outlined"
                    onChange={this._handleTextFieldChange}
                />

            </Grid>
            <Grid item xs={12} aligin="center">
                <Button variant="contained" color="primary" onClick={ this._roomButtonPressed }>
                    进入歌房
                </Button>
            
                <Button variant="contained" color="secondary" to="/" component={Link}>
                    回到首页
                </Button>
            </Grid>
        </Grid>
    }

    _handleTextFieldChange(e) {
        this.setState({
            roomCode: e.target.value,
        });
    };

    _roomButtonPressed() {
        console.log(this.state.roomCode);
        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: JSON.stringify({
                code: this.state.roomCode,
            }),
            mode: 'cors',
            credentials: "include",
        };
        fetch("http://192.168.199.133:9898/api/v1/Join_room", requestOptions).then((response) => 
            response.json()
        ).then((data) => {
            if (data.code === 200) {
                this.props.history.push(`/room/${this.state.roomCode}`)
            } else {
                this.setState({
                    error: "没有找到对应的歌房"
                })
            }
        }).catch((error) => {
            console.log(error);
        })
    };
}