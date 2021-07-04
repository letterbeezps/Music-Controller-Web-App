import React, { Component } from "react";
import Button from "@material-ui/core/Button";
import Grid from "@material-ui/core/Grid";
import Typography from "@material-ui/core/Typography";
import TextField from "@material-ui/core/TextField";
import FormHelperText from "@material-ui/core/FormHelperText";
import FormControl from "@material-ui/core/FormControl";
import { Link } from "react-router-dom";
import Radio from "@material-ui/core/Radio";
import RadioGroup from "@material-ui/core/RadioGroup";
import FormControlLabel from "@material-ui/core/FormControlLabel";
import { Collapse } from "@material-ui/core";
import Alert from "@material-ui/lab/Alert";

export default class CreateRoomPage extends Component {
    static defaultProps = {
        votesToSkip: 2,
        guestCanPause: true,
        update: false,
        roomCode: null,
        updateCallback: () => {},
    };

    constructor(props) {
        super(props);
        this.state = {
            guestCanPause: this.props.guestCanPause,
            votesToSkip: this.props.votesToSkip,
            errorMsg: "",
            successMsg: "",
        };

        this.handleRoomButtonPressed = this.handleRoomButtonPressed.bind(this)
        this.handleVotesChange = this.handleVotesChange.bind(this)
        this.handleGuestCanPauseChange = this.handleGuestCanPauseChange.bind(this)
        this.handleUpdateButtonPressed = this.handleUpdateButtonPressed.bind(this)
    }

    handleVotesChange(e) {
        this.setState({
            votesToSkip: e.target.value,
        });
    }

    handleGuestCanPauseChange(e) {
        this.setState({
            guestCanPause: e.target.value === "true" ? true : false,
        });
    }

    handleRoomButtonPressed() {
        console.log(this.state);
        let body = JSON.stringify({
            votes_to_skip: parseInt(this.state.votesToSkip),
            guest_can_pause: this.state.guestCanPause,
        })
        console.log(body)
        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: body,
            mode: 'cors',
            credentials: "include",
        };
        fetch("http://192.168.199.133:9898/api/v1/rooms", requestOptions).then((response) => 
            response.json()
        ).then((data) => this.props.history.push("/room/" + data.data.code));
    }

    handleUpdateButtonPressed() {
        console.log(this.state);
        let body = JSON.stringify({
            votes_to_skip: parseInt(this.state.votesToSkip),
            guest_can_pause: this.state.guestCanPause,
            code: this.props.roomCode,
        });
        console.log(body)
        const requestOptions = {
            method: 'POST',
            headers: {'Content-Type': 'application/json'},
            body: body,
            mode: 'cors',
            credentials: "include",
        };
        fetch("http://192.168.199.133:9898/api/v1/Update_room", requestOptions).then((response) => 
            response.json()
        ).then((data) => {
            if (data.code === 200) {
                this.setState({
                    successMsg: "更新成功"
                });
            } else {
                this.setState({
                    errorMsg: "更新失败"
                });
            }
            this.props.updateCallback();
        }); 
        
    }

    renderCreateButtons() {
        return (
            <Grid container spacing={1}>
                <Grid item xs={12} align="center">
                    <Button color="primary" variant="contained" onClick={this.handleRoomButtonPressed}>
                        创建歌房
                    </Button>
                    <Button color="secondary" variant="contained" to="/" component={Link}>
                        后退
                    </Button>
                </Grid>
            </Grid>
        );
    }

    renderUpdateButtons() {
        return (
            <Grid item xs={12} align="center">
                <Button
                    color="primary"
                    variant="contained"
                    onClick={this.handleUpdateButtonPressed}
                >
                    更新歌房
                </Button>
            </Grid>
        )
    }

    render() {
        const title = this.props.update ? "更新歌房" : "创建歌房"
        return <Grid container spacing={1}>
            <Grid item xs={12} align="center">
                <Collapse in={this.state.errorMsg !== "" || this.state.successMsg !== ""}>
                    {this.state.successMsg !== "" 
                        ? (
                            <Alert 
                                security="success" 
                                onClose={() => {
                                    this.setState({
                                        successMsg: ""
                                    });
                                }}
                                >
                                    {this.state.successMsg}
                            </Alert>
                          ) 
                        : (
                            <Alert 
                                severity="error"
                                onClose={() => {
                                    this.setState({
                                        errorMsg: ""
                                    });
                                }}
                                >
                                {this.state.errorMsg}
                            </Alert>
                          )
                    }
                </Collapse>
            </Grid>
            <Grid item xs={12} align="center">
                <Typography component='h4' variant='h4'>
                    {title}
                </Typography>
            </Grid>
            <Grid item xs={12} align="center">
                <FormControl component="fieldset">
                    <FormHelperText>
                        <div align="center">
                            访客是否可以控制播放状态
                        </div>
                    </FormHelperText>
                    <RadioGroup 
                        row 
                        defaultValue={this.props.guestCanPause.toString()} 
                        onChange={this.handleGuestCanPauseChange}
                    >
                        <FormControlLabel value="true" 
                            control={<Radio color="primary" />}
                            label="Play/Pause"
                            labelPlacement="bottom"
                        />

                        <FormControlLabel value="false" 
                            control={<Radio color="secondary" />}
                            label="No control"
                            labelPlacement="bottom"
                        />
                    </RadioGroup>
                </FormControl>
            </Grid>
            <Grid item xs={12} align="center">
                <FormControl>
                    <TextField 
                        required={true} 
                        type="number" 
                        onChange={this.handleVotesChange}
                        defaultValue={this.state.votesToSkip} 
                        inputProps={{
                            min: 1,
                            style: {textAlign: "center"},
                        }}
                    />
                    <FormHelperText>
                        <div align="center">
                            投票人数
                        </div>
                    </FormHelperText>
                </FormControl>
            </Grid> 
            {this.props.update 
                ? this.renderUpdateButtons() 
                : this.renderCreateButtons()}
        </Grid>
    }
}