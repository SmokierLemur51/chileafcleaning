from quart import Quart 
from pathlib import Path
from sqlite3 import dbapi2 as sqlite3
from quart_sqlalchemy import SQLAlchemy



def create_app(mode='Development'):
	""" in production you should use create_app('Production') """
	app = Quart(__name__, static_url_path='/static')
	# app.config.from_object(f"config.{mode}")

	from qblog.public.routes import public
	app.register_blueprint(public)

	from qblog.director.routes import director
	app.register_blueprint(director)

	return app

#____________________________________________________
# * * * * * * * * * * * * * * * * * * * * * * * * * * 
# * * * * * * * * * create app * * * * * * * * * * * *


app = create_app("Development")

app.config.update({
  "DATABASE": Path(app.root_path) / "blog.db",
})

#____________________________________________________
# * * * * * * * * * * * * * * * * * * * * * * * * * * 
# * * * * * * * * * connect db * * * * * * * * * * * *

def _connect_db():
    engine = sqlite3.connect(app.config["DATABASE"])
    engine.row_factory = sqlite3.Row
    return engine

def init_db():
    db = _connect_db()
    with open(Path(app.root_path) / "schema.sql", mode="r") as file_:
        db.cursor().executescript(file_.read())
    db.commit()

def _get_db():
	if not hasattr(g, "sqlite_db"):
		g.sqlite_db = _connect_db()
	return g.sqlite_db

#____________________________________________________
# * * * * * * * * * * * * * * * * * * * * * * * * * * 
# * * * * * * * * * run app * * * * * * * * * * * *

def run() -> None:
	app.run(debug=True)


