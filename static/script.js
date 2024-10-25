async function checkPermission() {
  const distributor = document.getElementById("distributor").value;
  const region = document.getElementById("region").value;
  const resultDiv = document.getElementById("result");
  const resultText = document.getElementById("resultText");

  if (!region) {
    alert("Please enter a region code");
    return;
  }

  try {
    const response = await fetch("/check-permission", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        distributorName: distributor,
        region: region,
      }),
    });

    if (!response.ok) {
      throw new Error("Network response was not ok");
    }

    const data = await response.json();

    resultDiv.classList.remove("hidden", "allowed", "denied");
    resultDiv.classList.add(data.hasPermission ? "allowed" : "denied");

    resultText.textContent = data.hasPermission
      ? `${distributor} is ALLOWED to distribute in ${region}`
      : `${distributor} is NOT ALLOWED to distribute in ${region}`;
  } catch (error) {
    console.error("Error:", error);
    alert("An error occurred while checking the permission");
  }
}
