resource "random_string" "random" {
  length = 10
}

module "test" {
  source = "./.."

  name = random_string.random.result
}