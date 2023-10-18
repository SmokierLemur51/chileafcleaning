from quart import Blueprint, render_template, g, redirect, request, url_for, current_app
from pathlib import Path
from sqlite3 import dbapi2 as sqlite3

public = Blueprint("public", __name__)


def _connect_db():
    engine = sqlite3.connect(current_app.config["DATABASE"])
    engine.row_factory = sqlite3.Row
    return engine


def _get_db():
	if not hasattr(g, "sqlite_db"):
		g.sqlite_db = _connect_db()
	return g.sqlite_db



@public.route("/")
async def index():
	return await render_template(f"public/index.html", title="Public Index")


@public.route("/services")
async def services():
	return await render_template(f"public/services.html", title="Services")


@public.route("/create-contact-request/", methods=["POST"])
async def create():
	if request.method == "POST":
		db = _get_db()
		form = await request.form
		db.execute(
			"INSERT INTO post (title, text) VALUES (?, ?)",
			[form["title"], form["text"]],
		)
		db.commit()
		return redirect(url_for("public.posts"))
	else:
		return redirect(url_for("public.posts"))


@public.route("/about")
async def about():
	db = _get_db()
	cur = db.execute(
		"""SELECT title, text
		   FROM post 
		ORDER BY id DESC""",
	)
	posts = cur.fetchall()
	return await render_template(f"public/posts.html", title="Posts", posts=posts)


