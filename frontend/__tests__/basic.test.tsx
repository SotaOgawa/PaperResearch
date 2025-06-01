import { render, screen } from "@testing-library/react";

function DummyComponent() {
  return <div>テスト表示</div>;
}

test("コンポーネントが表示される", () => {
  render(<DummyComponent />);
  expect(screen.getByText("テスト表示")).toBeInTheDocument();
});