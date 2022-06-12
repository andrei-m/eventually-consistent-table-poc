
function RandomAuto() {
  let handleSubmit = async (e) => {
    e.preventDefault();
    try {
      let res = await fetch("/api/randomized_auto", {
        method: "POST"
      });
      console.log(res.status);
    } catch (err) {
      console.log(err);
    }
  };

  return (
    <div>
      <form onSubmit={handleSubmit}>
        <button type="submit">Create Random Auto</button>
      </form>
    </div>
  );
}

export default RandomAuto;
