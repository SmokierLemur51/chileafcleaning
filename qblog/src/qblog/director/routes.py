from quart import Blueprint, render_template, request, redirect, url_for, g, current_app
from sqlite3 import dbapi2 as sqlite3
# from pathlib import Path


director = Blueprint("director", __name__, url_prefix="/director")

# database ... 
def _connect_db():
    engine = sqlite3.connect(current_app.config["DATABASE"])
    engine.row_factory = sqlite3.Row
    return engine


def _get_db():
	if not hasattr(g, "sqlite_db"):
		g.sqlite_db = _connect_db()
	return g.sqlite_db


# temp customer obj
class Customer:
	def __init__(self, name, idee):
		self.Name = name
		self.Id = idee


customers = [
	Customer("Logan", 1),
	Customer("Callie", 2),
	Customer("Steve", 3),
]


# routes 
@director.route("/")
async def portal():
	return await render_template(f"director/portal.html", title="Directors Corner")


#____________________________________________________
# * * * * * * * * * * * * * * * * * * * * * * * * * * 
# * * * * * * * * * tickets * * * * * * * * * * * *

@director.route("/tickets")
async def tickets():
	db = _get_db()
	cur = db.execute(
		"SELECT title, hourly_rate,"
	)
	return await render_template(f"director/tickets.html", title="Tickets", customers=customers)


@director.route("/create-ticket", methods=["POST"])
async def create_ticket():
	if request.method == "POST":
		db = _get_db()
		form = await request.form
		# check for values
		
		db.execute(
			"INSERT INTO ticket (title, hourly_rate, total_hours, description, client_id) VALUES (?, ?, ?, ?, ?, ?)",
			[
				form["ticket_title"], form["hourly_rate"], form["total_hours"], 
				form["description"], form["customer"],
			],
		)
		return redirect(url_for("director.tickets"))


#____________________________________________________
# * * * * * * * * * * * * * * * * * * * * * * * * * * 
# * * * * * * * * * clients * * * * * * * * * * * *
@director.route("/clients")
async def clients():
	
	return await render_template("clients.html", clients=clients, notes=notes, tickets=tickets)