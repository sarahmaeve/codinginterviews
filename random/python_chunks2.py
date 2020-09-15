# store all of the poll results
class Poll:
    # put options at init?
    def __init__(self, title, options):
        self.title = title
        self.results = {}
        for option in options:
            self.results[option] = 0

    def user_vote(self, user, option):
        # validate selection
        # if self.results[option]
        if self.title not in user.voted_in:
            try:
                self.results[option] += 1
                user.voted_in.append(self.title)
            except KeyError:
                print("That is not a valid selection in this poll.")
        else:
            print("User has already voted in that poll.")

        return

    def print_results(self):
        for k, v in self.results.items():
            print(k, ":", v, sep=" ")



# are we authenticating users?
class User:
    # self.user_name
    # voted_in array with poll
    def __init__(self, name):
        self.user_name = name
        self.voted_in = []

    # possible : list name / polls voted_in
    def print_name(self):
        print(self.user_name)

    def print_polls_voted(self):

        for poll in self.voted_in:
            print(poll)


p1_options = ["Green", "Blue", "Yellow", "Other"]
p1 = Poll("Your Favorite Color",p1_options)
# p1.print_results()

u1 = User("Sarah")
u1.print_name()
p1.user_vote(u1, "Blue")
p1.print_results()
# p1.user_vote(u1, "Yellow")
p1.print_results()

u2 = User("Alouette")
p1.user_vote(u2, "Other")
p1.print_results()
