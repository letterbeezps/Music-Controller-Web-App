/* eslint-disable no-useless-constructor */
import React, { Component } from "react";
import { Grid, Typography, Card, IconButton, LinearProgress } from "@material-ui/core";
import PlayArrowTcon from "@material-ui/icons/PlayArrow";
import SkipNextIcon from "@material-ui/icons/SkipNext";
import PauseIcon from "@material-ui/icons/Pause";

export default class MusicPlayer extends Component {
    constructor(props) {
        super(props);
    }

    pauseSong() {
        const requestOptions = {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            mode: 'cors',
            credentials: "include",
        }
        fetch("http://192.168.199.133:9898/api/spotify/pause", requestOptions);
    }

    playSong() {
        const requestOptions = {
            method: 'PUT',
            headers: {'Content-Type': 'application/json'},
            mode: 'cors',
            credentials: "include",
        }
        fetch("http://192.168.199.133:9898/api/spotify/play", requestOptions);
    }

    render() {
        const songProgress = (this.props.time / this.props.duration) * 100;
        return (
        <Card>
            <Grid container alignItems="center">
                <Grid item align="center" xs={4}>
                    <img src={this.props.image_url} height="100%" width="100%" alt="" />
                </Grid>
                <Grid item align="center" xs={8}>
                    <Typography component="h5" variant="h5">
                        {this.props.title}
                    </Typography>
                    <Typography color="textSecondary" variant="subtitle1">
                        {this.props.artist}
                    </Typography>
                    <div>
                        <IconButton onClick={ () => { 
                            this.props.is_playing ? this.pauseSong() : this.playSong(); 
                            }}
                        >
                            {this.props.is_playing ? <PauseIcon /> : <PlayArrowTcon />}
                        </IconButton>
                        <IconButton>
                            <SkipNextIcon />
                        </IconButton>
                    </div>
                </Grid>
            </Grid>
            <LinearProgress variant="determinate" value={ songProgress } />
        </Card>
        );
    }
}

