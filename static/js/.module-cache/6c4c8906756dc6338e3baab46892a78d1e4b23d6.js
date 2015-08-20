(function() {
    var Challenge = React.createClass({displayName: "Challenge",
        render: function() {
            var solved;
            if(this.props.challenge.solved) {
                solved = React.createElement("span", {className: "solved pull-right text-success ion-checkmark"}, " Solved")
            }

            return (
                React.createElement("div", {className: "col-sm-4"}, 
                    React.createElement("div", {className: "js-challenge challenge card"}, 
                        React.createElement("div", {className: "content content-border"}, 
                            React.createElement("span", {className: "title"}, this.props.challenge.id, ". ", this.props.challenge.name, " ", solved), 
                            React.createElement("p", null, this.props.challenge.smallDescription)
                        ), 
                        React.createElement("div", {className: "action"}, 
                            React.createElement("span", {className: "pull-left"}, "Solved By: ", this.props.challenge.solvedTotal), 
                            React.createElement("a", {href: "/problems/" + this.props.challenge.apiUrl}, "Solve")
                        )
                    )
                )
            );
        }
    });

    var ChallengeList = React.createClass({displayName: "ChallengeList",
        getInitialState: function() {
            return {
                showSolved: true,
                hiddenIds: []
            };
        },

        componentDidMount: function() {
            var ctx = this;
            var searchInput = document.getElementById('js-filter-challenges-input');
            searchInput.addEventListener('input', function() {
                ctx.onSearchInputChange(searchInput.value);
            });

            var solvedCheckbox = document.getElementById('js-filter-solved-checkbox');
            solvedCheckbox.addEventListener('click', function() {
                ctx.onSolvedCheckboxClick(solvedCheckbox.checked);
            });
        },

        render: function() {
            var ctx = this;
            var challenges = this.props.challenges.filter(function (challenge) {
                return ctx.state.hiddenIds.indexOf(challenge.id) == -1;
            });

            if(!this.state.showSolved) {
                challenges = challenges.filter(function(challenge) {
                    return challenge.solved == false;
                });
            }

            var rows = [];
            for(var i = 0; i < challenges.length; i+= 3) {
                var challengesComponents = [];
                for(var j = i; j < i + 3; j++) {
                    if(j < challenges.length) {
                        challengesComponents.push(
                            React.createElement(Challenge, {key: challenges[j].id, challenge: challenges[j]})
                        );
                    }
                }
                rows.push(
                    React.createElement("div", {className: "row challenge-row"}, 
                        challengesComponents
                    )
                );
            }

            return (
                React.createElement("div", {className: "container"}, 
                    rows
                )
            );
        },

        searchInString: function(str, search) {
            return str.toLowerCase().search(search.toLowerCase()) == -1 ? false : true;
        },

        onSearchInputChange: function(searchValue) {
            var hiddenIds = this.props.challenges
                                .filter(function(challenge) {
                                    return !this.searchInString(challenge.name, searchValue) && !this.searchInString(challenge.smallDescription, searchValue);
                                }, this)
                                .map(function(challenge) {
                                    return challenge.id;
                                });

            this.setState({
                hiddenIds: hiddenIds
            });
        },

        onSolvedCheckboxClick: function(checked) {
            this.setState({showSolved: checked});
        }
    });

    React.render(
      React.createElement(ChallengeList, {challenges: window.challenges}),
      document.getElementById('js-content')
    );
})();
