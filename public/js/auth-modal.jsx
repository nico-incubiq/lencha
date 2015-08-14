
(function() {

    var ModalAuth = React.createClass({

        getInitialState: function() {
            return {
                visible: false,
                error: false,
                isLoging: false,
                isRegistering: false
            };
        },

        componentDidMount: function() {
            var ctx = this;
            document.getElementById('js-login-register').onclick = function() {
                ctx.setState({
                    visible: true
                });
            };
        },

        onShow: function() {
            this.setState({
                visible: true
            });
        },

        onHide: function() {
            this.setState({
                visible: false
            });
        },

        handleSubmitLogin: function(e) {
            e.preventDefault();
            var username = React.findDOMNode(this.refs.loginUsername).value;
            var password = React.findDOMNode(this.refs.loginPassword).value;

            this.setState({isLoging: true, error: false});

            if(username === '' || password === '') {
                this.setState({isLoging: false, error: true, errorMessage: 'Sign in fields are empty !'});
                return;
            }

            var ctx = this;
            superagent
                .post('/login')
                .type('form')
                .send({username: username, password: password})
                .end(function(err, res) {
                    ctx.setState({isLoging: false});
                    if(err) {
                        if(err.status != 401) {
                            ctx.setState({error: true, errorMessage: 'Server error, try again please.'});
                        } else {
                            ctx.setState({error: true, errorMessage: res.body.message});
                        }
                    } else {
                        window.location.replace("/");
                    }
                });
        },

        handleSubmitRegister: function(e) {
            e.preventDefault();
            var username             = React.findDOMNode(this.refs.registerUsername).value;
            var password             = React.findDOMNode(this.refs.registerPassword).value;
            var passwordVerification = React.findDOMNode(this.refs.registerPasswordVerification).value;
            var email                = React.findDOMNode(this.refs.registerEmail).value;
            var emailUpdate          = React.findDOMNode(this.refs.registerEmailUpdate).checked;

            if(password != passwordVerification) {
                this.setState({isRegistering: false, error: true, errorMessage: "Password Mismatch"});
                return;
            }

            this.setState({isRegistering: true, error: false});

            var ctx = this;
            superagent
                .post('/register')
                .type('form')
                .send({
                    username: username,
                    password: password,
                    email: email,
                    emailUpdate: emailUpdate
                })
                .end(function(err, res) {
                    ctx.setState({isRegistering: false});
                    if(err) {
                        ctx.setState({error: true, success: false});
                        if(err.status != 401) {
                            ctx.setState({errorMessage: 'Server error, try again please.'});
                        } else {
                            ctx.setState({errorMessage: res.body.message});
                        }
                    } else {
                        ctx.setState({error: false, success: true, successMessage: res.body.message});
                    }
                });
        },

        render: function () {

            var loginBtn;
            if(!this.state.isLoging) {
                loginBtn = <button className="btn-signin btn btn-primary btn-block" type="submit">Sign in</button>;
            } else {
                loginBtn = (
                    <button className="btn-signin btn btn-primary btn-block disabled" type="submit" disabled="disabled">
                        <span className="ion-load-c"></span> Sign in
                    </button>
                );
            }

            var registerBtn;
            if(!this.state.isRegistering) {
                registerBtn = <button className="btn-signin btn btn-primary btn-block" type="submit">Register</button>;
            } else {
                registerBtn = (
                    <button className="btn-signin btn btn-primary btn-block disabled" type="submit" disabled="disabled">
                        <span className="ion-load-c"></span> Register
                    </button>
                );
            }

            var error;
            if(this.state.error) {
                error = (
                    <div className="alert-container">
                        <div className="alert alert-danger">{this.state.errorMessage}</div>
                    </div>
                );
            }

            var success;
            if(this.state.success) {
                error = (
                    <div className="alert-container">
                        <div className="alert alert-success">{this.state.successMessage}</div>
                    </div>
                );
            }

            return (
                <Modal visible={this.state.visible} closable={true} onShow={this.onShow} onHide={this.onHide}>
                    <h4 className="modal-title">Sign in or Create an account</h4>
                    {error}
                    {success}
                    <div className="row">
                        <div className="col-sm-6">
                            <form onSubmit={this.handleSubmitLogin} className="form-signin">
                                <h2>Sign in</h2>
                                <label htmlFor="inputEmail" className="sr-only">Email address</label>
                                <input name="username" type="text" className="form-control" placeholder="Username" required="" autofocus="" ref="loginUsername"/>
                                <label htmlFor="inputPassword" className="sr-only">Password</label>
                                <input name="password" type="password" className="form-control" placeholder="Password" required="" ref="loginPassword" />

                                {loginBtn}
                            </form>
                        </div>
                        <div className="col-sm-6">
                            <form onSubmit={this.handleSubmitRegister} className="form-signin">
                                    <h2>Create an account</h2>
                                    <label className="sr-only">Email address</label>
                                    <input name="username" type="text" className="form-control" placeholder="Username" required="" autofocus="" ref="registerUsername"/>

                                    <label className="sr-only">Password</label>
                                    <input name="password" type="password" className="form-control" placeholder="Password" required="" ref="registerPassword"/>

                                    <label className="sr-only">Password Verification</label>
                                    <input name="password-rep" type="password" className="form-control" placeholder="Verify password" required="" ref="registerPasswordVerification"/>

                                    <label className="sr-only">Email (optional)</label>
                                    <input name="email" type="email" className="form-control" placeholder="Email (optional)" ref="registerEmail"/>

                                    <div className="checkbox register-checkbox">
                                        <label>
                                            <input type="checkbox" ref="registerEmailUpdate"/> Send me an email when there is a new problem
                                        </label>
                                    </div>

                                    {registerBtn}
                            </form>
                        </div>
                    </div>
                </Modal>
            );
        }
    });

    React.render(<ModalAuth />, document.getElementById('auth-modal'));
})();
