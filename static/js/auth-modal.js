
(function() {

    var ModalAuth = React.createClass({displayName: "ModalAuth",

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
                loginBtn = React.createElement("button", {className: "btn-signin btn btn-primary btn-block", type: "submit"}, "Sign in");
            } else {
                loginBtn = (
                    React.createElement("button", {className: "btn-signin btn btn-primary btn-block disabled", type: "submit", disabled: "disabled"}, 
                        React.createElement("span", {className: "ion-load-c"}), " Sign in"
                    )
                );
            }

            var registerBtn;
            if(!this.state.isRegistering) {
                registerBtn = React.createElement("button", {className: "btn-signin btn btn-primary btn-block", type: "submit"}, "Register");
            } else {
                registerBtn = (
                    React.createElement("button", {className: "btn-signin btn btn-primary btn-block disabled", type: "submit", disabled: "disabled"}, 
                        React.createElement("span", {className: "ion-load-c"}), " Register"
                    )
                );
            }

            var error;
            if(this.state.error) {
                error = (
                    React.createElement("div", {className: "alert-container"}, 
                        React.createElement("div", {className: "alert alert-danger"}, this.state.errorMessage)
                    )
                );
            }

            var success;
            if(this.state.success) {
                error = (
                    React.createElement("div", {className: "alert-container"}, 
                        React.createElement("div", {className: "alert alert-success"}, this.state.successMessage)
                    )
                );
            }

            return (
                React.createElement(Modal, {visible: this.state.visible, closable: true, onShow: this.onShow, onHide: this.onHide}, 
                    React.createElement("h4", {className: "modal-title"}, "Sign in or Create an account"), 
                    error, 
                    success, 
                    React.createElement("div", {className: "row"}, 
                        React.createElement("div", {className: "col-sm-6"}, 
                            React.createElement("form", {onSubmit: this.handleSubmitLogin, className: "form-signin"}, 
                                React.createElement("h2", null, "Sign in"), 
                                React.createElement("label", {htmlFor: "inputEmail", className: "sr-only"}, "Email address"), 
                                React.createElement("input", {name: "username", type: "text", className: "form-control", placeholder: "Username", required: "", autofocus: "", ref: "loginUsername"}), 
                                React.createElement("label", {htmlFor: "inputPassword", className: "sr-only"}, "Password"), 
                                React.createElement("input", {name: "password", type: "password", className: "form-control", placeholder: "Password", required: "", ref: "loginPassword"}), 

                                loginBtn
                            )
                        ), 
                        React.createElement("div", {className: "col-sm-6"}, 
                            React.createElement("form", {onSubmit: this.handleSubmitRegister, className: "form-signin"}, 
                                    React.createElement("h2", null, "Create an account"), 
                                    React.createElement("label", {className: "sr-only"}, "Email address"), 
                                    React.createElement("input", {name: "username", type: "text", className: "form-control", placeholder: "Username", required: "", autofocus: "", ref: "registerUsername"}), 

                                    React.createElement("label", {className: "sr-only"}, "Password"), 
                                    React.createElement("input", {name: "password", type: "password", className: "form-control", placeholder: "Password", required: "", ref: "registerPassword"}), 

                                    React.createElement("label", {className: "sr-only"}, "Password Verification"), 
                                    React.createElement("input", {name: "password-rep", type: "password", className: "form-control", placeholder: "Verify password", required: "", ref: "registerPasswordVerification"}), 

                                    React.createElement("label", {className: "sr-only"}, "Email (optional)"), 
                                    React.createElement("input", {name: "email", type: "email", className: "form-control", placeholder: "Email (optional)", ref: "registerEmail"}), 

                                    React.createElement("div", {className: "checkbox register-checkbox"}, 
                                        React.createElement("label", null, 
                                            React.createElement("input", {type: "checkbox", ref: "registerEmailUpdate"}), " Send me an email when there is a new problem"
                                        )
                                    ), 

                                    registerBtn
                            )
                        )
                    )
                )
            );
        }
    });

    React.render(React.createElement(ModalAuth, null), document.getElementById('auth-modal'));
})();
