


const fetchForm = document.querySelector('.fetchForm');
const btn = document.querySelector('.btn');

const userName = "pppp-fan"; // ユーザ名
const graphId = "popo-tsk-counter"; // グラフID
const pwd = "pppp20240312";



const POST_URL = `https://pixe.la/v1/users/${userName}/graphs/${graphId}`; // POST-URL
const postFetch = () => {
    let formData = new FormData(fetchForm);
    // 増減値を取得
    let incrementValue = formData.get('incrementValue');

    // 入力日を取得
    let selectedDate = formData.get('selectedDate');
    const dataBody = {
        date: selectedDate,
        quantity: incrementValue
    };
    fetch(POST_URL, {
        method: 'POST',
        headers: {
            'X-USER-TOKEN': pwd,
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(dataBody)
    }).then((response) => {
        if (!response.ok) {
            console.log('error!');
        }
        console.log('ok!');
        return response.json();
    }).then((data) => {
        console.log(data);
        // フォームの値をクリア
        fetchForm.reset();
    }).catch((error) => {
        console.log(error);
    });
};
// データ送信
btn.addEventListener('click', postFetch, false);